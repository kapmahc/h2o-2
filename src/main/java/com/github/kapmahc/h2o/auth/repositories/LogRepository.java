package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Log;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.logR")
public interface LogRepository extends CrudRepository<Log, Long> {
}
