package com.github.kapmahc.h2o.auth.resources;

import com.codahale.metrics.annotation.Timed;
import com.github.kapmahc.h2o.auth.models.Locale;
import com.github.kapmahc.h2o.auth.dao.LocaleDao;
import io.dropwizard.hibernate.UnitOfWork;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import java.util.List;
import java.util.Optional;

@Path("/locales")
@Produces(MediaType.APPLICATION_JSON)
public class LocaleResource {
    @GET
    @Timed
    @UnitOfWork
    public List<Locale> getLocale(@QueryParam("lang") Optional<String> lang) {
        return localeDao.findByLang(lang.orElse(java.util.Locale.US.toLanguageTag()));
    }

    @POST
    @Timed
    @UnitOfWork
    public void setLocale(@QueryParam("lang") Optional<String> lang) {

    }

    @DELETE
    @Timed
    @UnitOfWork
    public void deleteLocale(@QueryParam("lang") Optional<String> lang) {

    }


    @Inject
    private LocaleDao localeDao;

    public void setLocaleDao(LocaleDao localeDao) {
        this.localeDao = localeDao;
    }
}
