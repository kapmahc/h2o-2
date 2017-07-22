package com.github.kapmahc.h2o.forum.repositories;

import com.github.kapmahc.h2o.forum.models.Article;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("forum.articleR")
public interface ArticleRepository extends CrudRepository<Article, Long> {
}
