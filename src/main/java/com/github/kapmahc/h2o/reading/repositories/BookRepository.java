package com.github.kapmahc.h2o.reading.repositories;

import com.github.kapmahc.h2o.reading.models.Book;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository("reading.bookR")
public interface BookRepository extends CrudRepository<Book, Long> {
}
