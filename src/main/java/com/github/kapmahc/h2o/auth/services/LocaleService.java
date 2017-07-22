package com.github.kapmahc.h2o.auth.services;

import com.github.kapmahc.h2o.auth.models.Locale;
import com.github.kapmahc.h2o.auth.repositories.LocaleRepository;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service("auth.localeS")
public class LocaleService {
    public void set(String lang, String code, String message) {
        Locale item = localeRepository.findByLangAndCode(lang, code);
        if (item == null) {
            item = new Locale();
            item.setCode(code);
            item.setLang(lang);
        }
        item.setMessage(message);
        localeRepository.save(item);

    }

    public String get(String lang, String code) {
        Locale item = localeRepository.findByLangAndCode(lang, code);
        return item == null ? null : item.getMessage();
    }

    @Resource
    LocaleRepository localeRepository;
}
