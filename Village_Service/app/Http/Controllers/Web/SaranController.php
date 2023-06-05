<?php

namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;
use App\Models\Saran;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Validation\Rule;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;

class SaranController extends Controller
{
    public function index(Request $request)
    {
        $saran = Saran::all();
        return view('web.Saran.saran', compact('saran'));
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Saran  $saran
     * @return \Illuminate\Http\Response
     */
    public function show(Saran $saran)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Saran  $saran
     * @return \Illuminate\Http\Response
     */
    public function destroy(Saran $saran)
    {
        $saran->delete();
        // return response()->json([
        //     'status' => 'success',
        //     'message' => 'Saran deleted successfully',
        // ]);
    }
}
