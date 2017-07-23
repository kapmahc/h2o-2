package com.github.kapmahc.h2o.auth;

import com.github.kapmahc.h2o.auth.models.Role;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;

@Configuration
public class SecurityConfiguration extends WebSecurityConfigurerAdapter {
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .csrf().ignoringAntMatchers(
                "/druid/**",
                "/monitoring")

                .and().authorizeRequests().antMatchers(
                HttpMethod.GET,
                "/",
                "/install",
                "/users/sign-up",
                "/users/confirm",
                "/users/confirm/*",
                "/users/unlock",
                "/users/unlock/*",
                "/users/forgot-password",
                "/users/reset-password/*",
                "/leave-words/new").permitAll()

                .and().authorizeRequests().antMatchers(
                HttpMethod.POST,
                "/install",
                "/users/sign-up",
                "/users/confirm",
                "/users/unlock",
                "/users/forgot-password",
                "/users/reset-password",
                "/leave-words").permitAll()

                .and().authorizeRequests().antMatchers(
                "/druid/**",
                "/monitoring",
                "/admin/**").hasRole(Role.ADMIN)

                .and().logout().logoutUrl("/users/sign-out").logoutSuccessUrl("/").invalidateHttpSession(true)
                .and().authorizeRequests().anyRequest().authenticated()
                .and().formLogin().loginPage("/users/sign-in").permitAll();
    }
}
