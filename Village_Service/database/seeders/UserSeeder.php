<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Hash;

class UserSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $admin = User::create([
            'nama' => 'Yoas Hutapea',
            'nik' => '1212021711020002',
            'no_telp'=> '081373030589',
            'alamat' => 'Jln Tugu Raja Hutapea',
            'tempat_lahir' => 'Jakarta',
            'password' => Hash::make('yoas12345'),
        ]);

        $admin->assignRole('admin');

        User::create([
            'nama' => 'Warga 1',
            'nik' => '1212021711020003',
            'no_telp'=> '081373030589',
            'alamat' => 'Jln Tugu Raja Hutapea',
            'tempat_lahir' => 'Jakarta',
            'password' => Hash::make('wargasatu'),
        ]);

        User::create([
            'nama' => 'warga 2',
            'nik' => '1212021711020004',
            'no_telp'=> '081373030589',
            'alamat' => 'Jln Tugu Raja Hutapea',
            'tempat_lahir' => 'Jakarta',
            'password' => Hash::make('wargadua'),
        ]);
    }
}
