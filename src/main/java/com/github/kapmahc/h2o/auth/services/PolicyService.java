package com.github.kapmahc.h2o.auth.services;

import com.github.kapmahc.h2o.auth.models.Policy;
import com.github.kapmahc.h2o.auth.models.Role;
import com.github.kapmahc.h2o.auth.models.User;
import com.github.kapmahc.h2o.auth.repositories.PolicyRepository;
import com.github.kapmahc.h2o.auth.repositories.RoleRepository;
import com.github.kapmahc.h2o.auth.repositories.UserRepository;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.io.Serializable;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Date;

@Service("auth.policyS")
public class PolicyService {
    public <T extends Serializable> void allow(User user, String role, int years, int months, int days) {
        this.allow(user, role, null, null, years, months, days);
    }

    public <T extends Serializable> void allow(User user, String role, Class<T> rty, int years, int months, int days) {
        this.allow(user, role, rty, null, years, months, days);
    }

    public <T extends Serializable> void allow(User user, String role, Class<T> rty, Long rid, int years, int months, int days) {
        Role r = getRole(role, rty, rid);
        Policy p = policyRepository.findByUserAndRole(user, r);
        if (p == null) {
            p = new Policy();
            p.setUser(user);
            p.setRole(r);
            Date now = new Date();
            p.setStartUp(now);
            p.setShutDown(
                    Date.from(
                            LocalDateTime.
                                    ofInstant(now.toInstant(), ZoneId.systemDefault()).
                                    plusYears(years).
                                    plusMonths(months).
                                    plusDays(days).
                                    atZone(ZoneId.systemDefault()).
                                    toInstant()
                    )
            );
        }

    }

    public <T extends Serializable> void deny(User user, String role) {
        this.deny(user, role, null, null);
    }

    public <T extends Serializable> void deny(User user, String role, Class<T> rty) {
        this.deny(user, role, rty, null);
    }

    public <T extends Serializable> void deny(User user, String role, Class<T> rty, Long rid) {
        Policy p = policyRepository.findByUserAndRole(user, getRole(role, rty, rid));
        if (p != null) {
            policyRepository.delete(p);
        }
    }

    public <T extends Serializable> boolean can(User user, String role) {
        return this.can(user, role, null, null);
    }

    public <T extends Serializable> boolean can(User user, String role, Class<T> rty) {
        return this.can(user, role, rty, null);
    }

    public <T extends Serializable> boolean can(User user, String role, Class<T> rty, Long rid) {
        Policy p = policyRepository.findByUserAndRole(user, getRole(role, rty, rid));
        Date now = new Date();
        return p != null && now.after(p.getStartUp()) && now.before(p.getShutDown());
    }


    private <T extends Serializable> Role getRole(String name, Class<T> rty, Long rid) {
        String type = rty == null ? null : rty.getTypeName();
        Role r = roleRepository.findByNameAndResourceTypeAndResourceId(name, type, rid);
        if (r != null) {
            return r;
        }
        r = new Role();
        r.setName(name);
        r.setResourceType(type);
        r.setResourceId(rid);
        return roleRepository.save(r);
    }

    @Resource
    UserRepository userRepository;
    @Resource
    PolicyRepository policyRepository;
    @Resource
    RoleRepository roleRepository;
}
