<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Kegiatan extends Model
{
    use HasFactory;
    protected $primaryKey = "id";
    protected $table = 'kegiatan';
    protected $fillable = [
        'judul',
        'tempat',
        'tanggal_mulai',
        'tanggal_akhir',
        'deskripsi',
    ];
}
