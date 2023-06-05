<?php

namespace App\Http\Controllers\API;

use App\Http\Controllers\Controller;
use App\Models\Pengajuan;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Validator;

class PengajuanController extends Controller
{
    public function select(Request $request, $id = null)
    {
        if ($id) {
            $pengajuan = Pengajuan::findOrFail($id);

            // Check if the authenticated user owns the pengajuan record
            if ($pengajuan->user_id !== Auth::id()) {
                return response()->json([
                    'error' => 'unauthorized',
                    'message' => 'You are not authorized to access this pengajuan.',
                ], 403);
            }

            return response()->json([
                'data' => $pengajuan,
            ]);
        }

        $pengajuan = Pengajuan::where('user_id', Auth::id())->get();

        return response()->json([
            'data' => $pengajuan,
        ]);
    }

    public function create(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'jenis_pengajuan' => 'required',
            'deskripsi' => 'required',
            'file' => 'required',
        ]);

        if ($validator->fails()) {
            return response()->json([
                'error' => 'validation_failed',
                'message' => $validator->errors()->first(),
            ], 422);
        }

        $user = Auth::user();

        $pengajuan = new Pengajuan();
        $pengajuan->user_id = $user->id;
        $pengajuan->user_name = $user->nama;
        $pengajuan->jenis_pengajuan = $request->jenis_pengajuan;
        $pengajuan->deskripsi = $request->deskripsi;
        $pengajuan->file = $request->file;
        $pengajuan->save();

        return response()->json([
            'message' => 'Pengajuan created successfully.',
            'data' => $pengajuan,
        ], 201);
    }

    public function update(Request $request, $id)
    {
        $pengajuan = Pengajuan::findOrFail($id);

        // Check if the authenticated user owns the pengajuan record
        if ($pengajuan->user_id !== Auth::id()) {
            return response()->json([
                'error' => 'unauthorized',
                'message' => 'You are not authorized to update this pengajuan.',
            ], 403);
        }

        $validator = Validator::make($request->all(), [
            'jenis_pengajuan' => 'required',
            'deskripsi' => 'required',
            'file' => 'required',
        ]);

        if ($validator->fails()) {
            return response()->json([
                'error' => 'validation_failed',
                'message' => $validator->errors()->first(),
            ], 422);
        }

        $pengajuan->jenis_pengajuan = $request->jenis_pengajuan;
        $pengajuan->deskripsi = $request->deskripsi;
        $pengajuan->file = $request->file;
        $pengajuan->save();

        return response()->json([
            'message' => 'Pengajuan updated successfully.',
            'data' => $pengajuan,
        ]);
    }

    public function delete(Request $request, $id)
    {
        $pengajuan = Pengajuan::findOrFail($id);

        // Check if the authenticated user owns the pengajuan record
        if ($pengajuan->user_id !== Auth::id()) {
            return response()->json([
                'error' => 'unauthorized',
                'message' => 'You are not authorized to delete this pengajuan.',
            ], 403);
        }

        $pengajuan->delete();

        return response()->json([
            'message' => 'Pengajuan deleted successfully.',
        ]);
    }
}
