package com.github.kapmahc.h2o.auth.resources;

import com.codahale.metrics.annotation.Metered;
import com.codahale.metrics.annotation.Timed;
import com.github.kapmahc.h2o.auth.models.Locale;
import io.dropwizard.jersey.caching.CacheControl;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import java.util.Optional;
import java.util.concurrent.TimeUnit;

@Path("/locales")
@Produces(MediaType.APPLICATION_JSON)
public class LocaleResource {
    @GET
    @Timed
//    @Metered
//    @CacheControl(maxAge = 6, maxAgeUnit = TimeUnit.HOURS)
    public Locale getLocale(@QueryParam("code") Optional<String> code) {
        Locale l = new Locale();
        l.setCode(code.orElse(""));
        return l;
    }

}
