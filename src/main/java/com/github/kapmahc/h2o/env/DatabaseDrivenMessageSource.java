package com.github.kapmahc.h2o.env;

import com.github.kapmahc.h2o.auth.services.LocaleService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.ResourceLoaderAware;
import org.springframework.context.support.AbstractMessageSource;
import org.springframework.core.io.ResourceLoader;
import org.springframework.lang.Nullable;

import javax.annotation.Resource;
import java.text.MessageFormat;
import java.util.Locale;

public class DatabaseDrivenMessageSource extends AbstractMessageSource implements ResourceLoaderAware {
    @Nullable
    @Override
    protected String resolveCodeWithoutArguments(String code, Locale locale) {
        return getText(code, locale);
    }

    @Nullable
    @Override
    protected MessageFormat resolveCode(String code, Locale locale) {
        return createMessageFormat(getText(code, locale), locale);
    }

    @Override
    public void setResourceLoader(ResourceLoader resourceLoader) {

    }


    private String getText(String code, Locale locale) {
        String lang = locale.toLanguageTag();
        String text = localeService.get(lang, code);
        if (text != null) {
            return text;
        }

        logger.debug("not find in database {} {}", lang, code);
        return getParentMessageSource().getMessage(code, null, locale);

    }


    @Resource
    LocaleService localeService;
    private final static Logger logger = LoggerFactory.getLogger(DatabaseDrivenMessageSource.class);


}
