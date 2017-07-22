package com.github.kapmahc.h2o.ops.mail.repositories;

import com.github.kapmahc.h2o.ops.mail.models.Alias;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.mail.aliasR")
public interface AliasRepository extends JpaRepository<Alias, Long> {
}
