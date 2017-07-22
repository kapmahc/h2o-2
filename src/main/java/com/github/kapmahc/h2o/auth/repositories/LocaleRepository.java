package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Locale;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.localeR")
public interface LocaleRepository extends JpaRepository<Locale, Long> {
    Locale findByLangAndCode(String lang, String code);
}
