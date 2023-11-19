import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {environment} from "../environments/environment";


@Injectable()

export class IncomeServices{

    constructor(private http: HttpClient) {
    }

    url = environment.API_BASE_URL + '/income';

    createIncome(income: any) {
        return this.http.post(this.url, income);
    }

}
