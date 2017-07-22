package com.github.kapmahc.h2o.survey.repositories;

import com.github.kapmahc.h2o.survey.models.Form;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository("survey.formR")
public interface FormRepository extends JpaRepository<Form, Long> {
}
