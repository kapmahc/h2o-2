package com.github.kapmahc.h2o.env;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.MessageSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.support.ResourceBundleMessageSource;
import org.springframework.web.servlet.LocaleResolver;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;
import org.springframework.web.servlet.i18n.CookieLocaleResolver;
import org.springframework.web.servlet.i18n.LocaleChangeInterceptor;

import java.util.Locale;

@Configuration
public class LocaleConfig implements WebMvcConfigurer {

    private final static String LOCALE = "locale";

    @Bean
    public LocaleResolver localeResolver() {
        CookieLocaleResolver clr = new CookieLocaleResolver();
        clr.setDefaultLocale(Locale.US);
        clr.setCookieName(LOCALE);
        clr.setCookieMaxAge(Integer.MAX_VALUE);
        clr.setLanguageTagCompliant(true);
        return clr;
    }

    @Bean("propertiesMessageSource")
    public MessageSource propertiesMessageSource() {

        ResourceBundleMessageSource ms = new ResourceBundleMessageSource();
        ms.setBasename(basename);
        ms.setDefaultEncoding(encoding);
        ms.setCacheSeconds(cacheSecond);
        ms.setUseCodeAsDefaultMessage(true);
        ms.setFallbackToSystemLocale(fallback);

        return ms;
    }

    @Bean("messageSource")
    public MessageSource messageSource(@Qualifier("propertiesMessageSource") MessageSource parent) {
        DatabaseDrivenMessageSource ms = new DatabaseDrivenMessageSource();
        ms.setParentMessageSource(parent);
        return ms;
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        LocaleChangeInterceptor interceptor = new LocaleChangeInterceptor();
        interceptor.setParamName(LOCALE);
        interceptor.setLanguageTagCompliant(true);
        registry.addInterceptor(interceptor);
    }

    @Value("${spring.messages.basename}")
    String basename;
    @Value("${spring.messages.cache-seconds}")
    int cacheSecond;
    @Value("${spring.messages.encoding}")
    String encoding;
    @Value("${spring.messages.fallback-to-system-locale}")
    boolean fallback;


}
