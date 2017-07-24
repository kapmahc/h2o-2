package com.github.kapmahc.h2o.auth.dao;

import com.github.kapmahc.h2o.auth.models.Locale;

import java.util.List;

public interface LocaleDao {
    List<Locale> findByLang(String lang);
}
