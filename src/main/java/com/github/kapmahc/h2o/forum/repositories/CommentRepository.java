package com.github.kapmahc.h2o.forum.repositories;

import com.github.kapmahc.h2o.forum.models.Comment;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("forum.commentR")
public interface CommentRepository extends CrudRepository<Comment, Long> {
}
