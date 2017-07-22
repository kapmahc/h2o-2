package com.github.kapmahc.h2o.auth.services;

import com.github.kapmahc.h2o.auth.models.Setting;
import com.github.kapmahc.h2o.auth.repositories.SettingRepository;
import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.crypto.encrypt.Encryptors;
import org.springframework.security.crypto.keygen.KeyGenerators;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.Serializable;

@Service("auth.settingS")
public class SettingService {
    public <T extends Serializable> T get(String k, Class<T> clazz) {
        Setting item = settingRepository.findByKey(k);
        if (item == null) {
            return null;
        }
        String val;
        if (item.isEncrypt()) {
            val = Encryptors.delux(this.key, item.getKey().substring(0, salt)).decrypt(item.getVal().substring(salt));
        } else {
            val = item.getVal();
        }
        return gson.fromJson(val, clazz);
    }

    public <T extends Serializable> void set(String k, T v) {
        this.set(k, v, false);
    }

    public <T extends Serializable> void set(String k, T v, boolean f) {
        String val = gson.toJson(v);

        String salt = KeyGenerators.string().generateKey();

        if (f) {
            val = Encryptors.delux(this.key, salt).encrypt(salt);
        }
        Setting item = settingRepository.findByKey(k);
        if (item == null) {
            item = new Setting();
            item.setKey(k);
        }
        item.setEncrypt(f);
        item.setKey(salt + val);
        settingRepository.save(item);
    }

    @PostConstruct
    void init() {
        gson = new Gson();
    }

    private Gson gson;
    @Resource
    SettingRepository settingRepository;
    @Value("app.secrets.aes")
    String key;
    private final int salt = 16;
}
