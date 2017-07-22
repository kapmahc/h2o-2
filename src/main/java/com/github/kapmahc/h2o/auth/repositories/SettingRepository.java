package com.github.kapmahc.h2o.auth.repositories;

import com.github.kapmahc.h2o.auth.models.Setting;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("auth.settingR")
public interface SettingRepository extends JpaRepository<Setting, Long> {
    Setting findByKey(String k);
}
