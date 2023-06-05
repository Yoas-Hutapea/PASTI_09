<?php

namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;
use App\Models\Pengumuman;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Validation\Rule;

class PengumumanController extends Controller
{
    public function index(Request $request)
    {
        $pengumuman = Pengumuman::all();
        return view('web.Pengumuman.pengumuman', compact('pengumuman'));
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        return view('web.Pengumuman.create');
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
            'tanggal.required' => 'Tanggal harus diisi.',
            'tanggal.after_or_equal' => 'Tanggal harus setelah atau sama dengan tanggal sekarang.',
            'judul.required' => 'Judul harus diisi.',
            'deskripsi.required' => 'Deskripsi harus diisi.',
        ];

        $validator = Validator::make($request->all(), [
            'tanggal' => 'required|date|after_or_equal:today',
            'judul' => 'required',
            'deskripsi' => 'required',
        ], $messages);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()->first(),
            ]);
        }

        Pengumuman::create($request->all());

        return response()->json([
            'status' => 'success',
            'message' => 'Pengumuman created successfully',
        ]);
    }


    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Pengumuman  $pengumuman
     * @return \Illuminate\Http\Response
     */
    public function show(Pengumuman $pengumuman)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Pengumuman  $pengumuman
     * @return \Illuminate\Http\Response
     */
    public function edit(Pengumuman $pengumuman)
    {
        return view('web.Pengumuman.update', compact('pengumuman'));
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Pengumuman  $pengumuman
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Pengumuman $pengumuman)
    {
        $messages = [
            'tanggal.required' => 'Tanggal harus diisi.',
            'tanggal.after_or_equal' => 'Tanggal harus setelah atau sama dengan tanggal sekarang.',
            'judul.required' => 'Judul harus diisi.',
            'deskripsi.required' => 'Deskripsi harus diisi.',
        ];

        $validator = Validator::make($request->all(), [
            'tanggal' => 'required|date|after_or_equal:today',
            'judul' => 'required',
            'deskripsi' => 'required',
        ], $messages);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()->first(),
            ]);
        }

        $pengumuman->update($request->all());

        return response()->json([
            'status' => 'success',
            'message' => 'Pengumuman updated successfully',
        ]);

        return redirect()->route('pengumuman.index');
    }
    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Pengumuman  $pengumuman
     * @return \Illuminate\Http\Response
     */
    public function destroy(Pengumuman $pengumuman)
    {
        $pengumuman->delete();
        return response()->json([
            'status' => 'success',
            'message' => 'Pengumuman deleted successfully',
        ]);
    }
}
