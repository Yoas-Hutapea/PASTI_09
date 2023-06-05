<?php

use App\Http\Controllers\Web\PendudukController;
use App\Http\Controllers\Web\KegiatanController;
use App\Http\Controllers\Web\PerangkatDesaController;
use App\Http\Controllers\Web\PengumumanController;
use App\Http\Controllers\Web\DashboardController;
use App\Http\Controllers\Web\AuthController;
use App\Http\Controllers\Web\PdfController;
use App\Http\Controllers\Web\SaranController;
use App\Http\Controllers\Web\PengajuanController;
use Illuminate\Support\Facades\Route;



/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return redirect()->route('login');
});

Route::get('/login', [AuthController::class, 'index'])->name('login');
Route::post('/login', [AuthController::class, 'do_login'])->name('do_login');

Route::group(['middleware' => ['role:admin']], function () {
    Route::get('/dashboard', [DashboardController::class, 'index'])->name('dashboard');
    Route::post('/logout', [AuthController::class, 'do_logout'])->middleware('auth:sanctum')->name('logout');

    //Penduduk
    Route::get('/penduduk', [PendudukController::class, 'index'])->name('penduduk.index');
    Route::get('/penduduk/profile', [PendudukController::class, 'profile'])->name('profile');


    Route::get('/penduduk/create', [PendudukController::class, 'create'])->name('penduduk.create');
    Route::post('/penduduk/create', [PendudukController::class, 'store'])->name('create-penduduk');

    Route::get('/penduduk/{penduduk}', [PendudukController::class, 'edit'])->name('penduduk.update');
    Route::put('/penduduk/{penduduk}', [PendudukController::class, 'update'])->name('update-penduduk');

    Route::delete('/penduduk/{penduduk}', [PendudukController::class, 'destroy'])->name('delete-penduduk');

    //Perangkat Desa
    Route::get('/perangkat', [PerangkatDesaController::class, 'index'])->name('perangkat.index');

    Route::get('/perangkat/create', [PerangkatDesaController::class, 'create'])->name('perangkat.create');
    Route::post('/perangkat/create', [PerangkatDesaController::class, 'store'])->name('create-perangkat');

    Route::get('/perangkat/{perangkat}', [PerangkatDesaController::class, 'edit'])->name('perangkat.update');
    Route::put('/perangkat/{perangkat}', [PerangkatDesaController::class, 'update'])->name('update-perangkat');

    Route::delete('/perangkat/{perangkat}', [PerangkatDesaController::class, 'destroy'])->name('delete-perangkat');

    //Pengumuman
    Route::get('/pengumuman', [PengumumanController::class, 'index'])->name('pengumuman.index');

    Route::get('/pengumuman/create', [PengumumanController::class, 'create'])->name('pengumuman.create');
    Route::post('/pengumuman/create', [PengumumanController::class, 'store'])->name('create-pengumuman');

    Route::get('/pengumuman/{pengumuman}', [PengumumanController::class, 'edit'])->name('pengumuman.update');
    Route::put('/pengumuman/{pengumuman}', [PengumumanController::class, 'update'])->name('update-pengumuman');

    Route::delete('/pengumuman/{pengumuman}', [PengumumanController::class, 'destroy'])->name('delete-pengumuman');

    //Kegiatan
    Route::get('/kegiatan', [KegiatanController::class, 'index'])->name('kegiatan.index');

    Route::post('/kegiatan', [KegiatanController::class, 'store'])->name('create-kegiatan');

    Route::get('/kegiatan/{kegiatan}', [KegiatanController::class, 'edit'])->name('kegiatan.update');
    Route::put('/kegiatan/{kegiatan}', [KegiatanController::class, 'update'])->name('update-kegiatan');

    Route::delete('/kegiatan/{kegiatan}', [KegiatanController::class, 'destroy'])->name('delete-kegiatan');

    //pdf
    Route::get('/penduduk/export-pdf', [PdfController::class, 'generateUserPdf'])->name('user-pdf');
    Route::get('/perangkat/pdf', [PdfController::class, 'generatePerangkatPdf'])->name('perangkat-pdf');

    //saran
    Route::get('/saran', [SaranController::class, 'index'])->name('saran.index');

    //pengajuan
    Route::get('/pengajuan', [PengajuanController::class, 'index'])->name('pengajuan.index');

});
