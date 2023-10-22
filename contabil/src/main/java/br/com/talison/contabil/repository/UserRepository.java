package br.com.talison.contabil.repository;

import br.com.talison.contabil.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UserRepository extends MongoRepository<User, String> {

}
