@extends('layouts.master')
@section('body')
    <div id="remoteModelData" class="modal fade" role="dialog"></div>
    <div class="content-page">
        <div class="container-fluid">
            <div class="row">
                <div class="col-lg-12">
                    <div class="d-flex flex-wrap align-items-center justify-content-between my-schedule mb-4">
                        <div class="d-flex align-items-center justify-content-between">
                            <h4>Kalender Kegiatan</h4>
                        </div>
                        <div class="create-workform">
                            <button type="button"
                                class="btn btn-primary position-relative d-flex align-items-center justify-content-between"
                                data-toggle="modal" data-target="#exampleModal">
                                <svg xmlns="http://www.w3.org/2000/svg" class="mr-2" width="20" fill="none"
                                    viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                        d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                                </svg>
                                Tambah Kegiatan Baru
                            </button>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-lg-12">
                            <div class="card card-block card-stretch">
                                <div class="card-body">
                                    <div id="calendar1" class="calendar-s"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Modal -->
        <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-body">
                        <div class="d-flex justify-content-between align-items-center pb-3">
                            <h5 class="modal-title" id="exampleModalLabel">Add New Event</h5>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <form id="kegiatanForm" method="POST" action="{{ route('create-kegiatan') }}">
                            @csrf
                            <div class="row">
                                <div class="col-md-12 mb-3">
                                    <label for="judul"
                                        class="form-label font-weight-bold text-muted text-uppercase">Judul</label>
                                    <input type="text" class="form-control" id="judul"
                                        placeholder="Masukkan judul kegiatan" name="judul">
                                </div>
                                <div class="col-md-12 mb-3">
                                    <label for="tempat"
                                        class="form-label font-weight-bold text-muted text-uppercase">Tempat</label>
                                    <input type="text" class="form-control" id="tempat"
                                        placeholder="Masukkan tempat pelaksanaan" name="tempat">
                                </div>
                                <div class="col-md-12">
                                    <label for="tanggal"
                                        class="form-label font-weight-bold text-muted text-uppercase">Tanggal Mulai dan
                                        Akhir Kegiatan</label>
                                </div>
                                <div class="col-md-6 mb-3">
                                    <div class="input-group">
                                        <input type="datetime-local" class="form-control" id="tanggal_mulai"
                                            name="tanggal_mulai" placeholder="Start Date">
                                    </div>
                                </div>
                                <div class="col-md-6 mb-3">
                                    <div class="input-group">
                                        <input type="datetime-local" class="form-control" id="tanggal_akhir"
                                            name="tanggal_akhir" placeholder="End Date">
                                    </div>
                                </div>
                                <div class="col-md-12 mb-3">
                                    <label for="deskripsi"
                                        class="form-label font-weight-bold text-muted text-uppercase">Deskripsi</label>
                                    <textarea class="form-control" id="deskripsi" rows="2" name="deskripsi" placeholder="Masukkan Deskripsi"></textarea>
                                </div>
                            </div>
                            <div class="d-flex justify-content-end">
                                <button type="submit" class="btn btn-primary">Tambah Kegiatan</button>
                            </div>
                        </form>

                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
@section('scripts')
    <script src="{{ asset('assets/auth/js/backend-bundle.min.js') }}"></script>

    <script src="{{ asset('assets/auth/vendor/fullcalendar/core/main.js') }}"></script>
    <script src="{{ asset('assets/auth/vendor/fullcalendar/daygrid/main.js') }}"></script>
    <script src="{{ asset('assets/auth/vendor/fullcalendar/timegrid/main.js') }}"></script>
    <script src="{{ asset('assets/auth/vendor/fullcalendar/list/main.js') }}"></script>

    <!-- Flextree Javascript-->
    <script src="{{ asset('assets/auth/js/flex-tree.min.js') }}"></script>
    <script src="{{ asset('assets/auth/js/tree.js') }}"></script>

    <!-- Table Treeview JavaScript -->
    <script src="{{ 'assets/auth/js/table-treeview.js' }}"></script>

    <!-- SweetAlert JavaScript -->
    <script src="{{ asset('assets/auth/js/sweetalert.js') }}"></script>

    <!-- Vectoe Map JavaScript -->
    <script src="{{ asset('assets/auth/js/vector-map-custom.js') }}"></script>

    <!-- Chart Custom JavaScript -->
    <script src="{{ asset('assets/auth/js/customizer.js') }}"></script>

    <script src="{{ asset('assets/auth/vendor/Leaflet/leaflet.js') }}"></script>

    <script src="{{ asset('assets/auth/vendor/vanillajs-datepicker/dist/js/datepicker-full.js') }}"></script>

    <script src="{{ asset('assets/auth/js/charts/progressbar.js') }}"></script>

    <!-- Chart Custom JavaScript -->
    <script src="{{ asset('assets/auth/js/chart-custom.js') }}"></script>
    <script src="{{ asset('assets/auth/js/charts/01.js') }}"></script>
    <script src="{{ asset('assets/auth/js/charts/02.js') }}"></script>

    <!-- slider JavaScript -->
    <script src="{{ asset('assets/auth/js/slider.js') }}"></script>

    <!-- Emoji picker -->
    <script src="{{asset('assets/auth/vendor/emoji-picker-element/index.js')}}" type="module"></script>
    <!-- app JavaScript -->
    <script src="{{ asset('assets/auth/js/app.js') }}"></script>
    <script>
        (function($) {
            "use strict";

            $(document).ready(function() {
                $('.select2js').select2();
                $.ajaxSetup({
                    headers: {
                        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                    }
                });

                $(document).on('click', '.loadRemoteModel', function(e) {
                    e.preventDefault();
                    var url = $(this).attr('href');

                    if (url.indexOf('#') == 0) {
                        $(url).modal('open');
                    } else {
                        $.get(url, function(data) {
                            $('#remoteModelData').html(data);
                            $('#remoteModelData').modal();
                            $('form').validator();
                            $(".datepicker").flatpickr({
                                dateFormat: "d-m-Y"
                            });
                        });
                    }
                });

                $(document).on('click', '[data-form="ajax"]', function(f) {
                    $('form').validator('update');
                    f.preventDefault();
                    var current = $(this);
                    current.addClass('disabled');
                    var form = $(this).closest('form');
                    var url = form.attr('action');
                    var fd = new FormData(form[0]);

                    $.ajax({
                        type: "POST",
                        url: url,
                        data: fd, // serializes the form's elements.
                        success: function(e) {
                            if (e.status == true) {
                                if (e.event == "submited") {
                                    showMessage(e.message);
                                    $(".modal").modal('hide');
                                }
                                if (e.event == 'refresh') {
                                    // showMessage(e.message);
                                    window.location.reload();
                                }
                                if (e.event == "callback") {
                                    showMessage(e.message);
                                    $(".modal").modal('hide');
                                    location.reload();
                                }
                            }
                            if (e.status == false) {
                                if (e.event == 'validation') {
                                    errorMessage(e.message);
                                }
                            }
                        },
                        error: function(error) {

                        },
                        cache: false,
                        contentType: false,
                        processData: false,
                    });
                    f.preventDefault(); // avoid to execute the actual submit of the form.

                });

                $(document).ready(function() {

                    $(document).on('change', '.change_status', function() {

                        var status = $(this).prop('checked') == true ? 1 : 0;
                        console.log(status)
                        var id = $(this).attr('data-id');
                        var type = $(this).attr('data-type');
                        $.ajax({
                            type: "GET",
                            dataType: "json",
                            url: "https://templates.iqonic.design/datum/laravel/public/changeStatus",
                            data: {
                                'status': status,
                                'id': id,
                                'type': type
                            },
                            success: function(data) {
                                alert(data.message);
                            }
                        });
                    })
                })

                $(document).on('click', '[data-toggle="tabajax"]', function(e) {
                    e.preventDefault();
                    var selectDiv = this;
                    ajaxMethodCall(selectDiv);
                });

                function ajaxMethodCall(selectDiv) {

                    var $this = $(selectDiv),
                        loadurl = $this.attr('data-href'),
                        targ = $this.attr('data-target'),
                        id = selectDiv.id || '';

                    $.post(loadurl, function(data) {
                        $(targ).html(data);
                        $('form').append('<input type="hidden" name="active_tab" value="' + id +
                            '" />');
                    });

                    $this.tab('show');
                    return false;
                }

                $('form[data-toggle="validator"]').on('submit', function(e) {
                    window.setTimeout(function() {
                        var errors = $('.has-error')
                        if (errors.length) {
                            $('html, body').animate({
                                scrollTop: "0"
                            }, 500);
                            e.preventDefault()
                        }
                    }, 0);
                });
            });
        })(jQuery);
    </script>
    <script>
        function deletePenduduk(id) {
            $.ajax({
                url: '/penduduk/' + id,
                type: 'DELETE',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                success: function() {
                    // Handle success response here
                    window.location.href = "{{ route('penduduk.index') }}";
                    console.log('Record deleted successfully');
                },
                error: function() {
                    // Handle error response here
                    console.log('An error occurred while deleting the record');
                }
            });
        }
    </script>
    <script>
        function redirectToPendudukUpdate(pendudukId) {
            window.location.href = "{{ route('penduduk.update', ':id') }}".replace(':id', pendudukId);
        }
    </script>
    <script>
        var calendar1;
        if (jQuery('#calendar1').length) {
            var calendarEl = document.getElementById('calendar1');

            calendar1 = new FullCalendar.Calendar(calendarEl, {
                selectable: true,
                plugins: ["timeGrid", "dayGrid", "list", "interaction"],
                timeZone: "UTC",
                defaultView: "dayGridMonth",
                contentHeight: "auto",
                eventLimit: true,
                dayMaxEvents: 4,
                header: {
                    left: "prev,next today",
                    center: "title",
                    right: "dayGridMonth,timeGridWeek,timeGridDay,listWeek"
                },
                dateClick: function(info) {
                    $('#schedule-start-date').val(info.dateStr)
                    $('#schedule-end-date').val(info.dateStr)
                    $('#date-event').modal('show')
                },
                events: [
                    @foreach ($kegiatan as $kegiatan)
                        {
                            title: '{{ $kegiatan->judul }}',
                            place: '{{ $kegiatan->tempat }}',
                            start: '{{ $kegiatan->tanggal_mulai }}',
                            end: '{{ $kegiatan->tanggal_akhir }}',
                            description: '{{ $kegiatan->deskripsi }}'
                        },
                    @endforeach
                ]
            });
            calendar1.render();

            $(document).on("submit", "#submit-schedule", function(e) {
                e.preventDefault()
                const title = $(this).find('#judul').val()
                const place = $(this).find('#tempat').val()
                const startDate = moment(new Date($(this).find('#tanggal_mulai').val()), 'YYYY-MM-DD').format(
                    'YYYY-MM-DD') + 'T05:30:00.000Z'
                const endDate = moment(new Date($(this).find('#tanggal_akhir').val()), 'YYYY-MM-DD').format(
                    'YYYY-MM-DD') + 'T05:30:00.000Z'
                const description = $(this).find('#deskripsi').val()
                const event = {
                    title: title,
                    place: place,
                    start: startDate || '2020-12-22T02:30:00',
                    end: endDate || '2020-12-12T14:30:00',
                    description: description,
                }
                $(this).closest('#date-event').modal('hide')
                calendar1.addEvent(event)
            })
        }
    </script>
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11.1.4/dist/sweetalert2.all.min.js"></script>
    <script>
        // Attach event listener to the form submission
        $('#kegiatanForm').submit(function(event) {
            event.preventDefault(); // Prevent the form from submitting

            // Perform your Ajax validation here
            $.ajax({
                url: '{{ route('create-kegiatan') }}', // Replace with your route name
                type: 'POST',
                data: $(this).serialize(), // Serialize form data
                success: function(response) {
                    // Handle the response from the server

                    if (response.status === 'success') {
                        // Validation successful, show success message
                        Swal.fire('Success', response.message, 'success').then(function() {
                            // Redirect to the desired page after successful validation
                            window.location.href =
                            '{{ route('kegiatan.index') }}'; // Replace with your desired route
                        });
                    } else {
                        // Validation failed, show error message
                        Swal.fire('Validation Error', response.message, 'error');
                    }
                },
                error: function(xhr, status, error) {
                    // Handle any errors that occur during the Ajax request
                    console.log('Ajax request error:', error);
                }
            });
        });
    </script>
@endsection
