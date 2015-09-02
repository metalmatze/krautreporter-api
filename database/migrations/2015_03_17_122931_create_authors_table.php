<?php

use Illuminate\Database\Migrations\Migration;

class CreateAuthorsTable extends Migration
{

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('authors', function ($table) {
            $table->increments('id');
            $table->string('name');
            $table->string('title');
            $table->string('url');
            $table->string('image');
            $table->text('biography')->nullable();
            $table->text('socialmedia')->nullable();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('authors');
    }
}
