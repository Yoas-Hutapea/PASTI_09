<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Daftar Perangkat Desa</title>
</head>
<body>
    <h1>Daftar Perangkat Desa</h1>
    <table>
        <thead>
            <tr>
                <th>Nama</th>
                <th>Jabatan</th>
            </tr>
        </thead>
        <tbody>
            @foreach($perangkat as $perangkat)
            <tr>
                <td>{{ $perangkat->nama }}</td>
                <td>{{ $perangkat->jabatan }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
</body>
</html>
