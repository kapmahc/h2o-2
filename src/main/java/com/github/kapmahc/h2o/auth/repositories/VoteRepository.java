package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Vote;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.voteR")
public interface VoteRepository extends CrudRepository<Vote, Long> {
}
