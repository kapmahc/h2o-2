package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Policy;
import com.github.kapmahc.h2o.auth.models.Role;
import com.github.kapmahc.h2o.auth.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.policyR")
public interface PolicyRepository extends JpaRepository<Policy, Long> {
    Policy findByUserAndRole(User user, Role role);
}
