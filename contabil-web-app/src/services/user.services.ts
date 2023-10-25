import { Injectable } from "@angular/core"
import { HttpClient } from '@angular/common/http';
import { Observable } from "rxjs";
import { User } from "src/domains/User";

@Injectable()
export class userServices {

    url = `http://localhost:8080/user`;

    constructor(private http: HttpClient) {}

    getUsers(): Observable<User[]> {
        return this.http.get<User[]>(`${this.url}/all`);
    }

    login(user: string, password: string): Observable<User> {
        const mock:User =  new User();
        mock.nome = user;
        mock.password = password;
        return new Observable<User>(observer => {
            observer.next(mock);
        })
        //return this.http.post<User>(`${this.url}/login`, {user, password});
    }
}

