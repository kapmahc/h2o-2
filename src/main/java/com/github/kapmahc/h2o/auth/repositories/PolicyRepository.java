package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Policy;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.policyR")
public interface PolicyRepository extends CrudRepository<Policy, Long> {
}
