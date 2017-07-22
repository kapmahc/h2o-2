package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Locale;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.localeR")
public interface LocaleRepository extends CrudRepository<Locale, Long> {
}
