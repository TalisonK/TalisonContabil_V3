import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environments/environment";


@Injectable()
export class TotalsServices{

    url = environment.API_BASE_URL + '/totals';

    constructor(private http: HttpClient) {    }

    getTimeline(userId: string, year: string, month: string){
        return this.http.post(`${this.url}/timeline`, {userId, year, month});
    }

}
