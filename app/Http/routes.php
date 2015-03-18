<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It's a breeze. Simply tell Laravel the URIs it should respond to
| and give it the controller to call when that URI is requested.
|
*/

Route::get('/', function() {
    return 'hallo';
});

Route::get('authors', 'AuthorsController@index');
Route::get('authors/{id}', 'AuthorsController@show');

Route::get('articles', 'ArticlesController@index');
Route::get('articles/{id}', 'ArticlesController@show');
