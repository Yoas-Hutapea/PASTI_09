<?php

namespace App\Http\Controllers\web;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use PDF;
use App\Models\User;
use Illuminate\Support\Facades\App;
use App\Models\PerangkatDesa;

class PdfController extends Controller
{
    public function generateUserPdf()
    {

        $user = User::all();
        $view = view('web.Pdf.penduduk', compact('penduduk'))->render();
        $pdf = App::make('dompdf.wrapper');
        $pdf->loadHTML($view);
    }

    public function generatePerangkatPdf()
    {
        $perangkat = Perangkat::all();

        $pdf = PDF::loadView('pdf.perangkat', compact('perangkat'));
        return $pdf->stream();
    }

}
