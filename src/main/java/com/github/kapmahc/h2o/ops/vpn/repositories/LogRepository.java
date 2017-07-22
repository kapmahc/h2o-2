package com.github.kapmahc.h2o.ops.vpn.repositories;

import com.github.kapmahc.h2o.ops.vpn.models.Log;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.vpn.logR")
public interface LogRepository extends CrudRepository<Log, Long> {
}
