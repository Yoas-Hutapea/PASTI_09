<?php

namespace App\Http\Controllers\web;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Validation\Rule;
use App\Models\Kegiatan;


class KegiatanController extends Controller
{
    public function index(Request $request)
    {
        $kegiatan = Kegiatan::all();
        return view('web.Kegiatan.kegiatan', compact('kegiatan'));
    }


    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $messages = [
            'judul.required' => 'Judul kegiatan harus diisi.',
            'tempat.required' => 'Tempat kegiatan harus diisi.',
            'tanggal_mulai.required' => 'Tanggal mulai kegiatan harus diisi.',
            'tanggal_mulai.date' => 'Tanggal mulai kegiatan harus dalam format tanggal yang valid.',
            'tanggal_akhir.required' => 'Tanggal akhir kegiatan harus diisi.',
            'tanggal_akhir.date' => 'Tanggal akhir kegiatan harus dalam format tanggal yang valid.',
            'tanggal_akhir.after_or_equal' => 'Tanggal akhir kegiatan harus setelah atau sama dengan tanggal mulai kegiatan.',
            'deskripsi.required' => 'Deskripsi kegiatan harus diisi.',
        ];

        $validator = Validator::make($request->all(), [
            'judul' => 'required',
            'tempat' => 'required',
            'tanggal_mulai' => 'required|date',
            'tanggal_akhir' => 'required|date|after_or_equal:tanggal_mulai',
            'deskripsi' => 'required',
        ], $messages);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()->first(),
            ]);
        }

        Kegiatan::create($request->all());

        return response()->json([
            'status' => 'success',
            'message' => 'Kegiatan created successfully',
        ]);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Kegiatan  $kegiatan
     * @return \Illuminate\Http\Response
     */
    public function show(Kegiatan $kegiatan)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Kegiatan  $kegiatan
     * @return \Illuminate\Http\Response
     */
    public function edit(Kegiatan $kegiatan)
    {
        return view('web.Kegiatan.update', compact('kegiatan'));
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Kegiatan  $kegiatan
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Kegiatan $kegiatan)
{
    $validator = Validator::make($request->all(), [
        'judul' => 'required',
        'tempat' => 'required',
        'tanggal_mulai' => 'required',
        'tanggal_akhir' => 'required',
        'deskripsi' => 'required'
    ]);

    if ($validator->fails()) {
        return response()->json([
            'status' => 'error',
            'message' => $validator->errors()->first(),
        ]);
    }

    $kegiatan->update($request->all());

    return response()->json([
        'status' => 'success',
        'message' => 'Kegiatan updated successfully',
    ]);

    return redirect()->route('kegiatan.index');
}
    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Kegiatan  $kegiatan
     * @return \Illuminate\Http\Response
     */
    public function destroy(Kegiatan $kegiatan)
    {
        $kegiatan->delete();
        return response()->json([
            'status' => 'success',
            'message' => 'Kegiatan deleted successfully',
        ]);
    }
}
