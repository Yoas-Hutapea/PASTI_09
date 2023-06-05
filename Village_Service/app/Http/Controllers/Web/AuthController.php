<?php

namespace App\Http\Controllers\Web;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Spatie\Permission\Contracts\Role;
use Illuminate\Support\Facades\Http;


class AuthController extends Controller
{
    public function __construct()
    {
        $this->middleware('guest')->except('do_logout');
    }

    public function index()
    {
        return view('auth.login');
    }

    public function do_login(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'nik' => 'required|min:16|max:16',
            'password' => 'required|min:8',
        ]);

        if ($validator->fails()) {
            $errors = $validator->errors();
            if ($errors->has('nik')) {
                return response()->json([
                    'alert' => 'error',
                    'message' => $errors->first('nik'),
                ]);
            } else {
                return response()->json([
                    'alert' => 'error',
                    'message' => $errors->first('password'),
                ]);
            }
        }

        $response = Http::post('http://localhost:8081/login', [
            'nik' => $request->nik,
            'password' => $request->password,
        ]);

        if ($response->successful()) {
            $data = $response->json();

            if ($data['alert'] === 'valid') {
                // Authentication succeeded

                // Find the user in your local database
                $user = User::where('nik', $request->nik)->first();

                if ($user) {
                    if (Hash::check($request->password, $user->password)) {
                        // Perform the login
                        if (Auth::attempt(['nik' => $request->nik, 'password' => $request->password], $request->remember)) {
                            // Redirect to dashboard or return JSON response
                            return response()->json([
                                'alert' => 'valid',
                                'message' => 'Berhasil Login',
                            ]);
                        }
                    } else {
                        // Invalid password
                        return response()->json([
                            'alert' => 'error',
                            'message' => 'Maaf, Password Salah.',
                        ]);
                    }
                } else {
                    // User not found
                    return response()->json([
                        'alert' => 'error',
                        'message' => 'Maaf, nik Salah atau belum terdaftar.',
                    ]);
                }
            } else {
                // Authentication failed
                // Return error response
                return response()->json([
                    'alert' => 'error',
                    'message' => 'Maaf, nik Salah atau belum terdaftar.',
                ]);
            }
        } else {
            // Failed to make the API request
            // Return error response
            return response()->json([
                'alert' => 'error',
                'message' => 'Failed to connect to the authentication service.',
            ]);
        }
    }

    public function do_logout()
    {
        Auth::guard('web')->logout();
        return redirect('/');
    }
}
