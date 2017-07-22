package com.github.kapmahc.h2o.ops.mail.repositories;

import com.github.kapmahc.h2o.ops.mail.models.Domain;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("ops.mail.domainR")
public interface DomainRepository extends CrudRepository<Domain, Long> {
}
