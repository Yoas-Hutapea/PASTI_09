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
                <th>Jenis Pengajuan</th>
                <th>Deskripsi</th>
                <th>File</th>
            </tr>
        </thead>
        <tbody>
            @foreach($pengajuan as $index => $pengajuan)
            <tr>
                <td>{{ $index + 1 }}</td>
                <td>{{ $user->user_name }}</td>
                <td>{{ $user->jenis_pengajuan }}</td>
                <td>{{ $user->deskripsi }}</td>
                <td>{{ $user->file }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
</body>
</html>
