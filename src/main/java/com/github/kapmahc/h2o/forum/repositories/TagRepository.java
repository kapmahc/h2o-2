package com.github.kapmahc.h2o.forum.repositories;

import com.github.kapmahc.h2o.forum.models.Tag;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("forum.tagR")
public interface TagRepository extends CrudRepository<Tag, Long> {
}
