package com.github.kapmahc.h2o.auth.services;

import com.github.kapmahc.h2o.auth.models.Log;
import com.github.kapmahc.h2o.auth.models.User;
import com.github.kapmahc.h2o.auth.repositories.LogRepository;
import com.github.kapmahc.h2o.auth.repositories.UserRepository;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.UnsupportedEncodingException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

@Service("auth.userS")
public class UserService {
    public User signIn(String email, String password) {
        User u = userRepository.findByProviderTypeAndProviderId(User.Type.EMAIL, email);
        if (u != null) {
            if (passwordEncoder.matches(password, u.getPassword())) {
                return u;
            }
        }
        return null;
    }

    public User addUser(String name, String email, String password) throws UnsupportedEncodingException, NoSuchAlgorithmException {
        User u = new User();
        u.setName(name);
        u.setEmail(email);
        u.setPassword(passwordEncoder.encode(password));
        u.setProviderId(email);
        u.setProviderType(User.Type.EMAIL);

        //https://en.gravatar.com/site/implement/images/java/
        byte[] md5 = MessageDigest.getInstance("MD5").digest(email.getBytes("CP1252"));
        StringBuilder buf = new StringBuilder();
        for (byte b : md5) {
            buf.append(Integer.toHexString((b & 0xFF) | 0x100).substring(1, 3));
        }
        u.setLogo(buf.toString());

        return userRepository.save(u);
    }

    public void addLog(User user, String message) {
        this.addLog(user, Log.Type.INFO, message);
    }

    public void addLog(User user, Log.Type type, String message) {
        Log l = new Log();
        l.setUser(user);
        l.setType(type);
        l.setMessage(message);
        logRepository.save(l);
    }

    @PostConstruct
    void init() {
        passwordEncoder = new BCryptPasswordEncoder(12);
    }

    @Resource
    UserRepository userRepository;
    @Resource
    LogRepository logRepository;
    private PasswordEncoder passwordEncoder;
}
