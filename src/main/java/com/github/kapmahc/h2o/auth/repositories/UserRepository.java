package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.User;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.userR")
public interface UserRepository extends CrudRepository<User, Long> {
}
