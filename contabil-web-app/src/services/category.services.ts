import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";


@Injectable()
export class CategoryServices {

    url = `http://localhost:8080/category`;

    constructor(private http: HttpClient) {
    };

    getCategories() {
        return this.http.get(`${this.url}/all`);
    }
}



