package com.github.kapmahc.h2o.ops.mail.repositories;

import com.github.kapmahc.h2o.ops.mail.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.mail.userR")
public interface UserRepository extends JpaRepository<User, Long> {
}
