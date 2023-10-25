import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';
import { PRIMENG_IMPORTS } from './primeng-imports';


@NgModule({
    imports: [
        PRIMENG_IMPORTS
    ],
    providers: [],
    exports: [
        PRIMENG_IMPORTS,
    ],
    schemas: [
        CUSTOM_ELEMENTS_SCHEMA
      ]
})
export class PrimengModule { }