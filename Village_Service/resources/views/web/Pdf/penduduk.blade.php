<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Daftar Penduduk</title>
</head>
<body>
    <h1>Daftar Penduduk</h1>
    <table>
        <thead>
            <tr>
                <th>Nama</th>
                <th>NIK</th>
                <th>No. Telp</th>
                <th>Alamat</th>
                <th>Tempat Lahir</th>
                <th>Tanggal Lahir</th>
                <th>Usia</th>
                <th>Jenis Kelamin</th>
                <th>Pekerjaan</th>
                <th>Agama</th>
                <th>KK</th>
            </tr>
        </thead>
        <tbody>
            @foreach($user as $index => $user)
            <tr>
                <td>{{ $index + 1 }}</td>
                <td>{{ $user->nama }}</td>
                <td>{{ $user->nik }}</td>
                <td>{{ $user->no_telp }}</td>
                <td>{{ $user->alamat }}</td>
                <td>{{ $user->tempat_lahir }}</td>
                <td>{{ $user->tanggal_lahir }}</td>
                <td>{{ $user->usia }}</td>
                <td>{{ $user->jenis_kelamin }}</td>
                <td>{{ $user->pekerjaan }}</td>
                <td>{{ $user->agama }}</td>
                <td>{{ $user->kk }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
</body>
</html>
