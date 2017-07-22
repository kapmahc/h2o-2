package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.FriendLink;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.friendLinkR")
public interface FriendLinkRepository extends CrudRepository<FriendLink, Long> {
}
