package com.github.kapmahc.h2o.ops.vpn.repositories;

import com.github.kapmahc.h2o.ops.vpn.models.Log;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.vpn.logR")
public interface LogRepository extends JpaRepository<Log, Long> {
}
