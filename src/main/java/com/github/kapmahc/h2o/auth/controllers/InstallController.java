package com.github.kapmahc.h2o.auth.controllers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller("auth.installC")
public class InstallController {
    //@PreAuthorize("permitAll()")
    @RequestMapping(value = "/install", method = RequestMethod.GET)
    public String getInstall(Model model) {
        return "auth/install";
    }

    //@PreAuthorize("permitAll()")
    @RequestMapping(value = "/install", method = RequestMethod.POST)
    public void postInstall() {
    }

    private final static Logger logger = LoggerFactory.getLogger(InstallController.class);
}
