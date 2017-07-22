package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Role;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.roleR")
public interface RoleRepository extends JpaRepository<Role, Long> {
    Role findByNameAndResourceTypeAndResourceId(String name, String resourceType, Long resourceId);
}
