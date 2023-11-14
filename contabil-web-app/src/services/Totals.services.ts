import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";


@Injectable()
export class TotalsServices{

    url = 'http://localhost:8080/totals';

    constructor(private http: HttpClient) {    }

    getTimeline(userId: string, year: string, month: string){
        return this.http.post(`${this.url}/timeline`, {userId, year, month});
    }

}
