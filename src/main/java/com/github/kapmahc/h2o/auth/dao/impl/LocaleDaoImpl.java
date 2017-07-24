package com.github.kapmahc.h2o.auth.dao.impl;

import com.github.kapmahc.h2o.auth.models.Locale;
import com.github.kapmahc.h2o.auth.dao.LocaleDao;
import io.dropwizard.hibernate.AbstractDAO;
import org.hibernate.SessionFactory;
import org.hibernate.query.Query;

import java.util.ArrayList;
import java.util.List;

public class LocaleDaoImpl extends AbstractDAO<Locale> implements LocaleDao {
    /**
     * Creates a new DAO with a given session provider.
     *
     * @param sessionFactory a session provider
     */
    public LocaleDaoImpl(SessionFactory sessionFactory) {
        super(sessionFactory);
    }

    @Override
    public List<Locale> findByLang(String lang) {
        Query<Locale> query = query("SELECT l FROM Locale l WHERE l.lang = :lang");
        query.setParameter("lang", lang);
        return list(query);
    }
}
