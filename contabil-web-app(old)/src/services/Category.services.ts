import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environments/environment";


@Injectable()
export class CategoryServices {

    url = environment.API_BASE_URL + '/category';

    constructor(private http: HttpClient) {
    };

    getCategories() {
        return this.http.get(`${this.url}/all`);
    }
}



