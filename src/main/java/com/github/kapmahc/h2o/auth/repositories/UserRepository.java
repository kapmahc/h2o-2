package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.userR")
public interface UserRepository extends JpaRepository<User, Long> {
    User findByUid(String uid);

    User findByProviderTypeAndProviderId(User.Type providerType, String providerId);
}
