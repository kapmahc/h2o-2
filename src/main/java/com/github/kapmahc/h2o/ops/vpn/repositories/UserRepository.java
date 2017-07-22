package com.github.kapmahc.h2o.ops.vpn.repositories;

import com.github.kapmahc.h2o.ops.vpn.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.vpn.userR")
public interface UserRepository extends JpaRepository<User, Long> {
}
