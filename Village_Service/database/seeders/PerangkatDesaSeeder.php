<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class PerangkatDesaSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        Perangkat::create([
            'nama' => 'Kalvin Hutapea',
            'jabatan' => 'Kepala Desa'
        ]);

        Perangkat::create([
            'nama' => 'Yoas Hutapea',
            'jabata' => 'Sekretaris Desa'
        ]);
    }
}
