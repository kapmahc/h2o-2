package com.github.kapmahc.h2o.survey.repositories;

import com.github.kapmahc.h2o.survey.models.Form;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("survey.formR")
public interface FormRepository extends CrudRepository<Form, Long> {
}
