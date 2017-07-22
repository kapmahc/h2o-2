package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.FriendLink;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.friendLinkR")
public interface FriendLinkRepository extends JpaRepository<FriendLink, Long> {
}
