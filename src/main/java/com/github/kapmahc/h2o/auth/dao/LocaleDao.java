package com.github.kapmahc.h2o.auth.dao;

import com.github.kapmahc.h2o.auth.models.Locale;
import io.dropwizard.hibernate.AbstractDAO;
import org.hibernate.SessionFactory;

import javax.inject.Inject;

public class LocaleDao extends AbstractDAO<Locale> {
    public Locale get(Long id){
        return get(id);
    }
    public LocaleDao( SessionFactory sessionFactory) {
        super(sessionFactory);
    }
}
