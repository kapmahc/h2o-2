package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Link;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.LinkR")
public interface LinkRepository extends CrudRepository<Link, Long> {
}
